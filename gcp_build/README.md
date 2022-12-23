# GCP Build
Because GitHub Actions is increasingly flakey, the weekly rebuilds of ATO images are now run via Google Cloud Platform,
using Compute Engine Spot Instances.

GitHub Actions is still used to trigger and monitor the build, because it provides good logging and notifications
support, but the actual build is executed on GCP.

Every Saturday at 9:52 AM, GitHub Actions runs the `start` script, which creates a VM instance, gets it to run the
`gcp_run` script, and then deletes the VM again.
