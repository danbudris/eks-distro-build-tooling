package constants

const AMD64Arch = "AMD64"
const AwsRegion = "us-east-1"
const ControlPlaneInstanceProfile = "arn:aws:iam::125833916567:instance-profile/KopsControlPlaneBuildRole"
const EksDPostSubmitArtifactsBucket = "eks-d-postsubmit-artifacts"
const KopsNodeInstanceProfile = "arn:aws:iam::125833916567:instance-profile/KopsNodesBuildRole"
const KopsStateStoreBucket = "s3://testbuildstack-125833916-kopsbuildstatestorebucke-d4esen60nfrk"
const ImageRepo = "public.ecr.aws/h1r8a7l5"
const DockerConfig = "/home/prow/go/src/github.com/aws/eks-distro/.docker"
const TestRoleArn = "arn:aws:iam::125833916567:role/TestBuildRole"
const ArtifactDeploymentRoleArn = "arn:aws:iam::379412251201:role/ArtifactDeploymentRole"

const K8s120 = "1.20"
const K8s121 = "1.21"
const K8s122 = "1.22"
const K8s123 = "1.23"
const K8s124 = "1.24"