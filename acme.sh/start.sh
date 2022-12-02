# 1. create out folder
mkdir -p /acme.sh/out/$MY_DOMAIN

# 2. set notification hook to telegram (https://github.com/acmesh-official/acme.sh/wiki/notify)
acme.sh --set-notify --notify-hook telegram --notify-level 2 --notify-mode 1

# 3. enable acme.sh auto upgrade
acme.sh --upgrade --auto-upgrade

# 4. register account with zerossl
acme.sh --register-account -m tiny656@hotmail.com # <your email>

# 5. issue cert by dns_azure way
acme.sh --issue --dns dns_azure -d $MY_DOMAIN -d *.$MY_DOMAIN

# 6. save cert results into `out` folder
acme.sh --install-cert -d $MY_DOMAIN \
--cert-file      /acme.sh/out/$MY_DOMAIN/$MY_DOMAIN.cer \
--ca-file        /acme.sh/out/$MY_DOMAIN/ca.cer \
--key-file       /acme.sh/out/$MY_DOMAIN/$MY_DOMAIN.key \
--fullchain-file /acme.sh/out/$MY_DOMAIN/fullchain.cer
