import click
import boto3
import threading
from os import path
from jinja2 import Environment, FileSystemLoader

KILT_CFN = path.join(path.dirname(__file__), 'kilt.yaml')
KILT_ZIP = path.join(path.dirname(__file__), '..', 'kilt.zip')

assert path.exists(KILT_CFN), 'Could not find cloudformation defaults'
assert path.exists(KILT_ZIP), 'Could not find ../kilt.zip - did you build it?'


class CallbackProgress:
    def __init__(self, filename, label):
        self._size = path.getsize(filename)
        self._lock = threading.Lock()
        self._bar = click.progressbar(label=label, length=self._size)
        self._bar.render_progress()

    def __call__(self, bytes_amount):
        with self._lock:
            self._bar.update(bytes_amount)

    def done(self):
        self._bar.render_finish()


@click.command('kilt-cfn-installer')
@click.argument('macro_name')
@click.argument('path_to_kilt_definition', type=click.Path(exists=True))
@click.option('--region', '-r', help="override aws region")
@click.option('--opt-in/--opt-out', is_flag=True, help="Use opt-in mechanism instead of the default opt out")
@click.option('--s3-arn', help="Use a different bucket instead of the default one (kilt-${account_id}-${region})")
@click.option('--kilt-zip-name', default="kilt.zip", help="[Optional] Name of the file for the lambda code")
def main(macro_name, path_to_kilt_definition, region, opt_in, s3_arn, kilt_zip_name):
    click.echo("Getting AWS account and region...", nl=False)
    aws_account = boto3.client('sts').get_caller_identity().get('Account')
    aws_region = boto3.session.Session().region_name
    if region is not None:
        aws_region = region
    click.echo(click.style(f'{aws_account} in {aws_region}', fg='green'))
    click.echo("Getting S3 bucket...", nl=False)
    s3_bucket = get_s3_bucket(aws_account, aws_region)
    click.echo(click.style(s3_bucket.name, fg='green'))
    pb = CallbackProgress(KILT_ZIP, "Uploading Macro")
    s3_bucket.upload_file(KILT_ZIP, kilt_zip_name, Callback=pb)
    pb.done()
    pb = CallbackProgress(path_to_kilt_definition, f"Uploading Kilt - {macro_name}")
    s3_bucket.upload_file(path_to_kilt_definition, f'{macro_name}.kilt.cfg', Callback=pb)
    pb.done()
    env = Environment(
        loader=FileSystemLoader(searchpath=path.dirname(__file__))
    )
    template = env.get_template('kilt.yaml')
    output_text = template.render(
        macro_name=macro_name,
        bucket_name=s3_bucket.name,
        kilt_zip_name=kilt_zip_name,
    )
    cf = boto3.resource('cloudformation', region_name=aws_region)
    stack_name = f'KiltMacro{macro_name}'
    click.echo(f"Creating stack {stack_name}...", nl=False)
    stack = cf.create_stack(
        StackName=stack_name,
        TemplateBody=output_text,
        Capabilities=[
            'CAPABILITY_IAM'
        ]
    )
    click.echo(click.style("SUBMITTED", fg="yellow"))
    click.echo(f"Check your cloudwatch console for stack {stack_name}")


def get_s3_bucket(aws_account, aws_region):
    s3_bucket_name = f'kilt-{aws_account}-{aws_region}'
    s3 = boto3.resource('s3')
    s3.create_bucket(
        ACL='private',
        Bucket=s3_bucket_name
    )
    bucket = s3.Bucket(s3_bucket_name)
    bucket.wait_until_exists()
    return bucket


if __name__ == '__main__':
    main()
