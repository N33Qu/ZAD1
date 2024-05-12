# server.py

import http.server
import socketserver
from datetime import datetime
import sys

PORT = 8080

AUTHOR_NAME = "Mikolaj Stasz"

class MyHandler(http.server.SimpleHTTPRequestHandler):
    def do_GET(self):
        current_time = datetime.now().strftime("%Y-%m-%d %H:%M:%S")
        author_info = f"Autor serwera: {AUTHOR_NAME}"
        port_info = f"Port: {PORT}"
        log_message = f"{current_time} - {author_info} - {port_info} - IP klienta: {self.client_address[0]}"

        print(log_message, file=sys.stderr)
        self.send_response(200)
        self.send_header('Content-type', 'text/html')
        self.end_headers()

        ip_address = self.client_address[0]
        html_response = f"<html><head><title>Informacje</title></head><body><h1>Witaj!</h1><p>Twój adres IP: {ip_address}</p><p>Teraz jest {current_time} w twojej strefie czasowej.</p></body></html>"
        self.wfile.write(html_response.encode())

with socketserver.TCPServer(("", PORT), MyHandler) as httpd:
    httpd.serve_forever()
