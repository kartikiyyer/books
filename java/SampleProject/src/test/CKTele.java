package test;

import java.util.Arrays;

import com.google.gson.Gson;

class Data {
	String[] users;
	String[] cc;
	int[][] cc_to_users;
}


public class CKTele {
	public static void main(String[] args) {
		String data = "{"
				+ "\"users\": [\"John\", \"Alice\", \"Dave\", \"Jean\"],"
				+ "\"cc\": [\"American Express\", \"Chase\", \"Amazon Prime\"],"
				+ "\"cc_to_users\": [[0,0],[1,1],[1,2],[2,0]]"
				+ "}";
		Gson g = new Gson();
		Data d = g.fromJson(data, Data.class);
		
		int[][] cc_to_count = new int[d.cc.length][2];
		for(int j=0; j< d.cc.length; j++) {
			cc_to_count[j][0] = -1;
			cc_to_count[j][1] = 0;
		}
		
		for(int i=0; i< d.cc_to_users.length; i++) {
			for(int j=0; j< cc_to_count.length; j++) {
				if(cc_to_count[j][0] != -1) {
					if(cc_to_count[j][0] == d.cc_to_users[i][0]) {
						cc_to_count[j][1]++;
						break;
					}
				} else {
					cc_to_count[j][0] = d.cc_to_users[i][0];
					cc_to_count[j][1]++;
					break;
				}

			}
		}
		int x = 0;
		int y = 0;
		for(int j=0; j< cc_to_count.length; j++) {
			System.out.println( cc_to_count[j][0] + " " + cc_to_count[j][1]);
			if(y == 0) {
				y = cc_to_count[j][1];
				x = cc_to_count[j][0];
			} else if(y > cc_to_count[j][1]) {
				y = cc_to_count[j][1];
				x = cc_to_count[j][0];
			}
		}
		System.out.println();
		System.out.println("Lowest cc used");
		System.out.println(x + " " + y);
	}
}
