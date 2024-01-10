clear
#!/bin/bash
#by lilmoe
v='\e[1;32m'
printf "\n"
echo '#!/bin/sh' >> run
echo 'export PATH=.' >> run
echo "-bash" >>run
echo "SERVER 94.125.182.255 6667      # Budapest.HU.EU">>lilmoe.set
echo "SERVER 154.35.175.201 6667      # Chicago.IL.US">>lilmoe.set
echo "SERVER 208.64.123.210 6667      # LosAngeles.CA.US">>lilmoe.set
echo "SERVER ix.undernet.org 6667     # ix.undernet.org">>lilmoe.set
echo "SERVER 198.148.91.146 6667      # Denver.CO.US">>lilmoe.set
echo "SERVER us.undernet.org 6667     # US.Undernet">>lilmoe.set
echo "SERVER eu.undernet.org 6667     # EU.Undernet">>lilmoe.set
echo "SERVER 94.125.182.255 7000      # Budapest.HU.EU">>lilmoe.set
echo "SERVER 154.35.175.201 7000      # Chicago.IL.US">>lilmoe.set
echo "SERVER 208.64.123.210 7000      # LosAngeles.CA.US">>lilmoe.set
echo "SERVER 198.148.91.146 7000      # Denver.CO.US">>lilmoe.set

printf "\n${v}"
printf "              lilmoe BlackMech v1.0  \n\n"

printf "Introdu Nick Bot: "
read NICK
echo "NICK $NICK">>lilmoe.set

echo "USERFILE                usr">>lilmoe.set
echo "CMDCHAR                 .">>lilmoe.set

printf "Introdu Ident Bot: "
read LOGIN
echo "LOGIN $LOGIN">>lilmoe.set

printf "Introdu Real-Name Bot: "
read IRCNAME
echo "IRCNAME $IRCNAME">>lilmoe.set

printf "Introdu USER Bot: "
read USER
printf "Introdu PASS USER Bot: "
read PASS
echo "SERVICE                msg X@channels.undernet.org login $USER $PASS ">>lilmoe.set


echo "VIRTUAL"                       >>lilmoe.set
echo "MODES                   +ix-ws">>lilmoe.set
echo "HASONOTICE                    ">>lilmoe.set
echo "TOG CC                  1">>lilmoe.set
echo "TOG CLOAK               1">>lilmoe.set
echo "TOG SPY                 1">>lilmoe.set
echo "SET OPMODES             6">>lilmoe.set
echo "SET BANMODES            6">>lilmoe.set
echo "SET CTIMEOUT            60">>lilmoe.set
echo "SET CDELAY              30">>lilmoe.set

printf "Introdu Chan Bot (ex #lilmoe): "
read CHANNEL
echo "CHANNEL                 $CHANNEL">>lilmoe.set

echo "TOG PUB                 1">>lilmoe.set
echo "TOG MASS                1">>lilmoe.set
echo "TOG SHIT                1">>lilmoe.set
echo "TOG PROT                1">>lilmoe.set
echo "TOG ENFM                0">>lilmoe.set
echo "TOG RV                  1">>lilmoe.set
echo "SET AAWAY               1">>lilmoe.set
echo "SET MDL                 3">>lilmoe.set
echo "SET MKL                 3">>lilmoe.set
echo "SET MBL                 3">>lilmoe.set
echo "SET MPL                 2">>lilmoe.set

echo "#### BOT 2 ####">>lilmoe.set

printf "\n"
printf "Setari Bot 2\n"
printf "\n"

printf "Introdu Nick Bot2: "
read NICK
echo "NICK $NICK">>lilmoe.set

echo "USERFILE                usr2">>lilmoe.set
echo "CMDCHAR                 .">>lilmoe.set

printf "Introdu Ident Bot2: "
read LOGIN
echo "LOGIN $LOGIN">>lilmoe.set

printf "Introdu Real-Name Bot2: "
read IRCNAME
echo "IRCNAME $IRCNAME">>lilmoe.set

printf "Introdu USER Bot2: "
read USER
printf "Introdu PASS USER Bot2: "
read PASS
echo "SERVICE                msg X@channels.undernet.org login $USER $PASS ">>lilmoe.set


echo "VIRTUAL"                       >>lilmoe.set
echo "MODES                   +ix-ws">>lilmoe.set
echo "HASONOTICE                    ">>lilmoe.set
echo "TOG CC                  1">>lilmoe.set
echo "TOG CLOAK               1">>lilmoe.set
echo "TOG SPY                 1">>lilmoe.set
echo "SET OPMODES             6">>lilmoe.set
echo "SET BANMODES            6">>lilmoe.set
echo "SET CTIMEOUT            60">>lilmoe.set
echo "SET CDELAY              30">>lilmoe.set

printf "Introdu Chan Bot2 (ex #lilmoe): "
read CHANNEL
echo "CHANNEL                 $CHANNEL">>lilmoe.set

echo "TOG PUB                 1">>lilmoe.set
echo "TOG MASS                1">>lilmoe.set
echo "TOG SHIT                1">>lilmoe.set
echo "TOG PROT                1">>lilmoe.set
echo "TOG ENFM                0">>lilmoe.set
echo "TOG RV                  1">>lilmoe.set
echo "SET AAWAY               1">>lilmoe.set
echo "SET MDL                 3">>lilmoe.set
echo "SET MKL                 3">>lilmoe.set
echo "SET MBL                 3">>lilmoe.set
echo "SET MPL                 2">>lilmoe.set


printf "\n"
printf "Autorun..."
printf "\n"
chmod +x lilmoe.set
chmod +x a
./a

printf "\n"
printf "Pornire Emech:                  [DA]"
printf "\n\n"
chmod +x run
./run

