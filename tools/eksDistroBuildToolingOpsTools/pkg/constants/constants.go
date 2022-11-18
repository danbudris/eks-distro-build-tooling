package constants

const AMD64Arch = "AMD64"
const ARM64Arch = "ARM64"

const DefaultAwsRegion = "us-east-1"

const EksDPostSubmitArtifactsBucket = "eks-d-postsubmit-artifacts"
const EksDGitRepoUrl = "https://github.com/aws/eks-distro"
const EKsDBuildToolingImageRepo = "public.ecr.aws/h1r8a7l5"

const EksDArtifactDeploymentRoleArn = "arn:aws:iam::379412251201:role/ArtifactDeploymentRole"
const EksDTestRoleArn = "arn:aws:iam::125833916567:role/TestBuildRole"

const DockerConfigPath = "/home/prow/go/src/github.com/aws/eks-distro/.docker"
const EksDKopsNodeInstanceProfile = "arn:aws:iam::125833916567:instance-profile/KopsNodesBuildRole"
const EKsDKopsStateStoreBucket = "s3://testbuildstack-125833916-kopsbuildstatestorebucke-d4esen60nfrk"
const EksDKopsControlPlaneRole = "arn:aws:iam::125833916567:instance-profile/KopsControlPlaneBuildRole"
