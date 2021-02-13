set -ex

ZONE=us-central1-f

gcloud compute instances delete my-app-instance --zone=$ZONE --delete-disks=all

gcloud compute firewall-rules delete default-allow-http-80