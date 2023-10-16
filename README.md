# Kuber Opt Action

Github Action for operating remote kubernetes by using GRPC.

You should deploy [this pod](https://github.com/NoahAmethyst/simple-kube-operator) in your kubernetes first

**When would you use it?**

When you want to operate your kubernetes automatic and login to server is Inconvenient.


## Arguments

| Argument Name           | Required | Default             | Description                                                              |
|-------------------------|----------|---------------------|--------------------------------------------------------------------------|
| `server`                | True     | N/A                 | The GRPC server address.Make sure it is availiable out the internal net. |
| `action`                | True     | N/A                 | The action you want to operate.For now just [delete-pod].                |
| `app`                   | False    | N/A                 | The app witch set in label.app.                                          |
| `namespace`             | False    | default             | Kubernetes namespace.                                                    |
| `pod-name`              | False    | N/A                 | The name of the pod witch you want to operate.                           |


## Example

```yaml
      - name: Operate Kubernetes
        uses: NoahAmethyst/kuber-opt-action@v1.0.0
        with:
          server: ${{ secrets.GRPC_SERVER }}
          action: delete-pod
          app: ${{ secrets.APP }}
```
