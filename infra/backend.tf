terraform {
  backend "gcs" {
    bucket = "ikigai-tfstate"
    prefix = "tf/state"
  }
}