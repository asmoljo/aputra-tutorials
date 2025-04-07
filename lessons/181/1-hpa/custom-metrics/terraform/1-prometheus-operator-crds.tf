# Install Manually
# helm repo add prometheus-community https://prometheus-community.github.io/helm-charts
# helm repo update
# helm install prometheus-operator-crds --namespace monitoring --version 7.0.0 prometheus-community/prometheus-operator-crds
resource "helm_release" "prometheus_operator_crds" {
  name = "prometheus-operator-crds"

  repository       = "https://prometheus-community.github.io/helm-charts"
  chart            = "prometheus-operator-crds"
  namespace        = "monitoring"
  create_namespace = true
  version          = "7.0.0"
}
