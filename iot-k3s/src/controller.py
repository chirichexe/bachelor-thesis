import kopf
import kubernetes.client as k8s
import kubernetes.config as k8s_config
from kubernetes.client.rest import ApiException

print("Controller avviato con successo...")

# Carica la configurazione del cluster
try:
    k8s_config.load_incluster_config()  # Se in un Pod dentro Kubernetes
except Exception:
    k8s_config.load_kube_config()  # Se in locale

@kopf.on.create('iot.example.com', 'v1', 'iotdevices')
@kopf.on.update('iot.example.com', 'v1', 'iotdevices')
def assign_device(spec, meta, namespace, **kwargs):
    device_name = meta.get('name')
    device_status = spec.get('status')
    target_ns = spec.get('targetNamespace', None)
    assigned_ns = spec.get('assignedNamespace', None)

    if device_status == "available" and target_ns and not assigned_ns:
        patch = {
            "spec": {
                "status": "assigned",
                "assignedNamespace": target_ns
            }
        }

        pod_manifest = {
            "apiVersion": "v1",
            "kind": "Pod",
            "metadata": {
                "name": f"device-{device_name}-pod",
                "labels": {
                    "iot-device": device_name
                }
            },
            "spec": {
                "containers": [
                    {
                        "name": "device-agent",
                        "image": "busybox",
                        "command": ["sh", "-c", "while true; do echo Device running; sleep 10; done"]
                    }
                ]
            }
        }

        core_v1 = k8s.CoreV1Api()
        try:
            core_v1.create_namespaced_pod(namespace=target_ns, body=pod_manifest)
            kopf.info(meta, reason="PodCreated", message=f"Pod device-{device_name}-pod creato in {target_ns}")
        except ApiException as e:
            raise kopf.TemporaryError(f"Errore nella creazione del Pod: {e}", delay=30)

        return patch

    kopf.info(meta, reason="NoChange", message="Nessuna azione necessaria")
