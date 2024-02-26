apiVersion: serving.knative.dev/v1
kind: Service
metadata:
  name: ikigai
  labels:
    cloud.googleapis.com/location: $REGION
  annotations:
    run.googleapis.com/ingress: internal
spec:
  template:
    metadata:
      annotations:
        autoscaling.knative.dev/minScale: "0"
        autoscaling.knative.dev/maxScale: "3"
        run.googleapis.com/startup-cpu-boost: "false"
    spec:
      containerConcurrency: 30
      timeoutSeconds: 60
      serviceAccountName: $SERVICE_ACCOUNT_EMAIL
      containers:
        - image: $IMAGE
          ports:
            - containerPort: 8080
          env:
            - name: OPENWEATHER_API_KEY
              valueFrom:
                secretKeyRef:
                  name: weather_api_key
                  key: latest
          resources:
            limits:
              cpu: "1"
              memory: "128Mi"
          startupProbe:
            tcpSocket:
              port: 8080
            initialDelaySeconds: 0
            timeoutSeconds: 2
            failureThreshold: 3
            periodSeconds: 20
          livenessProbe:
            httpGet:
              path: "/health"
              port: 8080
            initialDelaySeconds: 0
            timeoutSeconds: 2
            failureThreshold: 3
            periodSeconds: 40