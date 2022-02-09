if __name__=='__main__':
	w, h = 6,6
	input = [[0 for x in range(w)] for y in range(h)]
	input[0] = [-9,-9,-9,1,1,1]
	input[1] = [0,-9,0,4,3,2]
	input[2] = [-9,-9,-9,1,2,3]
	input[3] = [0,0,8,6,6,0]
	input[4] = [0,0,0,-2,0,0]
	input[5] = [0,0,1,2,4,0]
	print input
	
	for x in range(0,4):
		for y in range(0,4):
			a = b = 0
			add = 0
			flag = True
			for i in range(x,x+3):
				if flag:
					flag = False
				else:
					flag = True
				for j in range(y,y+3):
					if flag:
						if a == b: 
							add += input[i][j]
					else:
						add += input[i][j]
					b +=1
				a +=1
				b = 0
			print add


			