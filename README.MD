# 🛡️ HashiCorp Vault Setup with File-Based Storage in Docker

This guide sets up HashiCorp Vault using file-based storage **inside the container** (no external volumes). It includes initialization, unsealing, login, and storing `.env` secrets.

---

## 📦 Step 1: Run Vault Container

```bash
docker run --cap-add=IPC_LOCK \
  -p 8800:8800 \
  --name vault-file \
  -e 'VAULT_LOCAL_CONFIG={
    "ui": true,
    "listener": [{
      "tcp": {
        "address": "0.0.0.0:8800",
        "tls_disable": true
      }
    }],
    "storage": {
      "file": {
        "path": "/vault/data"
      }
    }
  }' \
  hashicorp/vault:1.15.5

```
## Step 2: Initialize Vault

``` bash
docker exec -it vault-file vault operator init
```

⚠️ Save these securely. You’ll need 3 of the unseal keys to unlock the Vault after a restart.

---

## Step 3: Unseal Vault
Use 3 out of the 5 unseal keys returned in Step 2:

``` bash
docker exec -it vault-file vault operator unseal <key1>
docker exec -it vault-file vault operator unseal <key2>
docker exec -it vault-file vault operator unseal <key3>
```

## Step 4: Login to Vault
Set the Vault address and log in using the root token:

``` bash
export VAULT_ADDR=http://localhost:8800
vault login <root_token>
```

## 🗄️ Step 5: Store Secrets (.env-style)
``` bash
vault kv put secret/app \
DB_USER=admin \
DB_PASS=secret123 \
JWT_SECRET=my-super-secret
```

## 📥 Step 6: Retrieve Stored Secrets

```bash
vault kv get secret/app
```

You should see output like:

```pgsql
Key          Value
----         -----
DB_USER      admin
DB_PASS      secret123
JWT_SECRET   my-super-secret
```

## ✅ Step 7: Notes
- Vault UI: http://localhost:8800
- Secrets are stored in /vault/data inside the container
- Secrets persist across restarts, but:
- You must unseal again after a restart
- If you docker rm the container, secrets are lost
- This setup is suitable for local development
