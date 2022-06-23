#!/bin/bash

login=0
for i in `seq 3`; do
  code=$(awk 'BEGIN{srand();printf("%06d",1000000*rand())}')
  /janbar/sbin/sendmail /janbar/etc/sendmail.json "$code"
  if ! read -t 60 -s -p 'code: ' inputCode ;then
    exit 0
  fi
  echo

  if [ "$inputCode" = "$code" ];then
    login=1; break
  fi
  echo -e 'Login incorrect\n'
done

if [ $login -eq 1 ];then
  /bin/bash
fi
