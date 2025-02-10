import requests

from sign_encryptor import generate_md5_signature

# 接口请求域名
main_url = 'https://api.salesmartly.com'

# 项目API Token 具体获取可以参考以下文档：https://help.salesmartly.com/docs/apitoken?search=1
api_token = 'api_token'

# 项目Id （在系统内左下角的卡片处获取，所有请求必传）
project_id = 'project_id'

# 发送 GET 请求
def send_get_request(api_url,params=None, headers=None):
    """
    发送带有查询参数和自定义头部的 GET 请求。

    :param api_url: 请求的具体接口
    :param params: 字典形式的查询参数
    :param headers: 字典形式的请求头
    """
    url = main_url + api_url  # 示例 API URL

    try:
        # 发送 GET 请求，附加查询参数和头部
        response = requests.get(url, params=params, headers=headers)
        response.raise_for_status()  # 检查请求是否成功

        # 将响应转换为 JSON
        data = response.json()
        print("GET 请求成功，响应数据：")
        print(data)

    except requests.exceptions.RequestException as e:
        print(f"GET 请求失败：{e}")

# 发送 POST 请求
def send_post_request(api_url,payload=None, headers=None):
    """
    发送带有自定义头部的 POST 请求，使用表单数据。

    :param api_url: 请求的具体接口
    :param payload: 字典形式的查询参数
    :param headers: 字典形式的请求头
    """
    url = main_url + api_url  # 示例 API URL

    try:
        # 发送 POST 请求，使用表单数据
        response = requests.post(url, data=payload, headers=headers)
        response.raise_for_status()  # 检查请求是否成功

        # 将响应转换为 JSON
        data = response.json()
        print("POST 请求成功，响应数据：")
        print(data)

    except requests.exceptions.RequestException as e:
        print(f"POST 请求失败：{e}")

if __name__ == "__main__":
    # 定义请求的接口及参数
    get_url = '/api/chat-user/get-contact-list'
    get_params = {
        'project_id': project_id,
        'updated_time': '{"start":1680000000,"end":1814027206}'
    }

    # 参数加密出签名
    get_sign = generate_md5_signature(api_token, get_params)
    get_header = {
        'external-sign': get_sign
    }
    send_get_request(get_url, get_params, get_header)
