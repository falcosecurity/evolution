# == Class: falco::repo
class falco::repo inherits falco {
  case $::osfamily {
    'Debian': {
      include apt::update

      Apt::Source [ 'falco' ]
      -> Class [ 'apt::update' ]

      apt::source { 'falco':
        location => 'https://download.falco.org/packages/deb',
        release  => 'stable',
        repos    => '',
        key      => {
          source => 'https://falco.org/repo/falcosecurity-3672BA8F.asc',
          id     => '3672BA8F'
        },
      }

      ensure_packages(["linux-headers-${::kernelrelease}"])
    }
    'RedHat': {
      include 'epel'

      Yumrepo [ 'falco' ]
      -> Class [ 'epel' ]

      yumrepo { 'falco':
        baseurl  => 'https://download.falco.org/packages/rpm',
        descr    => 'Falco repository',
        enabled  => 1,
        gpgcheck => 0,
      }

      ensure_packages(["kernel-devel-${::kernelrelease}"])
    }
    default: {
      fail("\"${module_name}\" provides no repository information for OSfamily \"${::osfamily}\"")
    }
  }
}
