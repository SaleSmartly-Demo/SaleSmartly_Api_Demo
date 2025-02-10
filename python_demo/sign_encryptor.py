import hashlib

def generate_md5_signature(api_token: str, data: dict) -> str:
    """
    根据字典序对字典进行排序，将排序后的键值对与 API token 组合，并生成 MD5 哈希。

    :param data: 要处理的字典
    :param api_token: 用于组合的 API token
    :return: 生成的 MD5 哈希值
    """
    # 将字典按键进行排序
    sorted_items = sorted(data.items())

    # 构建排序后的字符串，格式为 "key1=value1&key2=value2..."
    sorted_string = "&" + "&".join(f"{key}={value}" for key, value in sorted_items)

    # 将排序后的字符串与 API token 组合
    combined_string = api_token + sorted_string
    print("生成排序后的字符串：" + combined_string)

    # 生成 MD5 哈希
    md5_hash = hashlib.md5(combined_string.encode()).hexdigest()
    print("加密后的字符串：" + md5_hash)

    return md5_hash
