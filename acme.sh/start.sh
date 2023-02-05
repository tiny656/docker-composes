# 1. create out folder
mkdir -p /acme.sh/out/$MY_DOMAIN

# 2. set notification hook to telegram (https://github.com/acmesh-official/acme.sh/wiki/notify)
acme.sh --set-notify --notify-hook telegram --notify-level 2 --notify-mode 1

# 3. enable acme.sh auto upgrade
acme.sh --upgrade --auto-upgrade

# 4. register account with zerossl
acme.sh --register-account -m $MY_EMAIL

# 4. issue cert by dns_azure way, install cert results into `out` folder
# 4.1 issue cert by dns_azure way
acme.sh --issue --dns dns_azure -d $MY_DOMAIN -d *.$MY_DOMAIN -k ec-256 \
--renew-hook "acme.sh --install-cert -d $MY_DOMAIN \
--cert-file      /acme.sh/out/$MY_DOMAIN/$MY_DOMAIN.cer \
--ca-file        /acme.sh/out/$MY_DOMAIN/ca.cer \
--key-file       /acme.sh/out/$MY_DOMAIN/$MY_DOMAIN.key \
--fullchain-file /acme.sh/out/$MY_DOMAIN/fullchain.cer"
# 4.2 issue cert by standalone way
# acme.sh --issue -d $MY_DOMAIN --standalone -k ec-256 \
# --renew-hook "acme.sh --install-cert -d $MY_DOMAIN \
# --cert-file      /acme.sh/out/$MY_DOMAIN/$MY_DOMAIN.cer \
# --ca-file        /acme.sh/out/$MY_DOMAIN/ca.cer \
# --key-file       /acme.sh/out/$MY_DOMAIN/$MY_DOMAIN.key \
# --fullchain-file /acme.sh/out/$MY_DOMAIN/fullchain.cer"
# you can also try with letsencrypt, add `--server letsencrypt` behind
