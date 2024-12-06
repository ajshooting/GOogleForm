url = "https://docs.google.com/forms/d/e/1FAIpQLSfoEtgxNENKW3cK9Nfm-z7RGtEcUbdrrKxuyZkOveDUykgR-w/viewform"


import re
import requests
from pprint import pprint


def fetch(url):
    response = requests.get(url)
    return response.text


send_url = re.sub("viewfor.+", "formResponse", url)
print("url : " + send_url)

site_html = fetch(url)

# idリスト取得
# id_list = re.findall(r'\d,\[\[(\d*?),', site_html)
# id_list = [re.sub(r'\d,\[\[', "", i).replace(",", "") for i in id_list]
# id_per = list(set(id_list))

# 多分初期値
# name_list = re.findall(r'%\.@\.\[\d+,"(.*?)",', site_html)
# name_list = [re.sub(r'%\.@\.\[\d+,"', "", i).replace('",', "") for i in name_list]
# name_per = list(set(name_list))

id_list = []
q_list = []

# すべて取得
all_list = re.findall(r"%\.@\.\[(\d.*?>)", site_html)


for n in range(len(all_list)):
    print(all_list[n])

print("================")


# 取得
for n in range(len(all_list)):
    id_list.append(re.match(r"^\d+", all_list[n]).group())
    q_list.append(re.match(r"\d+,&quot;(.*?)&quot;,", all_list[n]).group(1))

pprint(q_list)
pprint(id_list)

# print(name_list)
# print(name_per)

# print("=====")
# dictionary = {key: '' for key in id_list}
# pprint(dictionary)

Data = []

# 配列の各要素に対して変数を作成
for i in range(len(q_list)):
    Data.append([q_list[i], id_list[i]])

print(Data)

payload = {}

for i in range(len(Data)):
    payload["entry." + Data[i][1]] = input(Data[i][0] + " : ")

pprint(payload)


for t in range(10000):
    response = requests.post(url, data=payload)
    print(str(t) + " : " + str(response.status_code))
