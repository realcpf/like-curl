package curl.v1;

import java.io.BufferedReader;
import java.io.IOException;
import java.io.InputStreamReader;
import java.io.PrintWriter;
import java.net.InetSocketAddress;
import java.net.Socket;
import java.sql.Timestamp;
import java.time.LocalDateTime;

//  https://curl.se/docs/httpscripting.html
public class Main {
	
	public static void main(String[] args) throws IOException {
		Socket client = new Socket();
		InetSocketAddress address = new InetSocketAddress("localhost",8000);
		
		System.out.println(LocalDateTime.now());
		client.connect(address,1000);
		System.out.println(LocalDateTime.now());
		String req = "GET /book/ HTTP/1.1\r\n" + 
		             "Host: localhost:8000\r\n\r\n";
		try(PrintWriter printWriter = new PrintWriter(client.getOutputStream(),true)){
			printWriter.println(req);
			String info = null;
			System.out.println(LocalDateTime.now());
			try(BufferedReader bReader = new BufferedReader(new InputStreamReader(client.getInputStream(),"UTF-8"))){
				System.out.println(LocalDateTime.now());
				while ((info = bReader.readLine())!=null) {
					System.out.println(info);
				}
			}
		}
	}

}
