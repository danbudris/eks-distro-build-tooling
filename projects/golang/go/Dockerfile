FROM public.ecr.aws/amazonlinux/amazonlinux:2 as builder
ARG ARTIFACTS_BUCKET
ARG GO_SOURCE_VERSION
ARG BUILD_ROOT
ARG MAKE_TARGET

RUN yum install make -y

COPY . ${BUILD_ROOT}

RUN make install-deps -C ${BUILD_ROOT}

RUN make ${MAKE_TARGET} -C ${BUILD_ROOT}


FROM scratch

COPY --from=builder ${BUILD_ROOT}/rpmbuild ./