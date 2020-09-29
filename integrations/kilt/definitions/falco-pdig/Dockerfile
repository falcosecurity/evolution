FROM alpine:latest as pdig-build

RUN apk add g++ gcc cmake cmake make libtool elfutils-dev libelf-static linux-headers git
RUN mkdir /source
RUN git clone https://github.com/falcosecurity/pdig /source/pdig
RUN git clone https://github.com/draios/sysdig /source/sysdig
RUN mkdir /source/pdig/build
RUN cd /source/pdig/build && cmake -DMUSL_OPTIMIZED_BUILD=True ..
RUN cd /source/pdig/build && make

FROM falcosecurity/falco:latest-slim AS falco

FROM scratch
COPY --from=falco /usr/bin/falco /falco/bin/falco
COPY --from=falco /usr/share/falco/lua /falco/share/lua
COPY --from=falco /etc /falco/etc
COPY --from=pdig-build /source/pdig/build/pdig /falco/bin/pdig
ADD binaries/waitforever /falco/bin/
ADD binaries/logshipper /falco/bin/
ADD binaries/launcher /falco/bin/

CMD ["/falco/bin/waitforever"]