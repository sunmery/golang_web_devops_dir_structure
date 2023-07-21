chcp 65001

set WG_CONFIG=C:\Program Files\WireGuard\lixia.conf
cd "C:\Program Files\WireGuard"
@REM start /b wg-quick.exe up %WG_CONFIG%
start /b wireguard.exe up %WG_CONFIG%
exit
