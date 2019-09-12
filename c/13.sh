curl -sSL https://projecteuler.net/problem=13|awk 'BEGIN{sum=0}{ if($0 ~ /[[:digit:]]+<br \/>/){ sum += substr($0,0,50)}; }END{print substr(sum,0,10)}'
