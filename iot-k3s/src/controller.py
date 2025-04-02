import kopf
import kubernetes.client as k8s
from kubernetes.client.rest import ApiException
import os

# Carica la configurazione quando il controller gira dentro il cluster
try:
    k8s.config.load_incluster_config()
except Exception:
    k8s.config.load_kube_config()

@kopf.on.create('iot.example.com', 'v1', 'iotdevices')
def assign_device(spec, meta, namespace, **kwargs):
    device_name = meta.get('name')
    device_status = spec.get('status')
    assigned_ns = spec.get('assignedNamespace')
    
    # Se non è stato specificato il namespace di assegnazione, usa quello corrente
    if not assigned_ns:
        assigned_ns = namespace

    # Se il dispositivo è disponibile, aggiorna lo stato e crea un Pod
    if device_status == "available":
        # Definisci la patch per aggiornare lo stato a "assigned" e registrare il namespace
        patch = {
            "spec": {
                "status": "assigned",
                "assignedNamespace": assigned_ns
            }
        }

        # Definisci il Pod associato al dispositivo
        pod_manifest = {
            "apiVersion": "v1",
            "kind": "Pod",
            "metadata": {
                "name": f"device-{device_name}-pod",
                "namespace": assigned_ns,
                "labels": {
                    "iot-device": device_name
                }
            },
            "spec": {
                "containers": [
                    {
                        "name": "device-agent",
                        "image": "busybox",  # Per test; in un caso reale potresti usare l'immagine del tuo agente IoT
                        "command": ["sh", "-c", "while true; do echo Device running; sleep 10; done"]
                    }
                ]
            }
        }

        # Crea il Pod nel namespace assegnato
        core_v1 = k8s.CoreV1Api()
        try:
            core_v1.create_namespaced_pod(namespace=assigned_ns, body=pod_manifest)
            kopf.info(meta, reason="PodCreated", message=f"Pod device-{device_name}-pod creato in {assigned_ns}")
        except ApiException as e:
            raise kopf.TemporaryError(f"Errore durante la creazione del Pod: {e}", delay=30)

        return patch
