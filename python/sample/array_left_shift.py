if __name__=='__main__':
	items = [1,2,3,4,5]
	for d in range(0,2):
		count = 0
		i = 0
		val_new = None
		while count < 5:
			if val_new is None:
				val = items[i]
			else:
				val = val_new
			if i == 0:
				i = 4
			else:
				i = i - 1
			val_new = items[i]
			items[i] = val
			count = count + 1
			print items

	print ''

	items = [1,2,3,4,5]
	for d in range(0,2):
		count = 0
		i = 0
		j = 4
		while count < 4:
			val = items[i]
			val_new = items[j]
			items[j] = val
			items[i] = val_new
			count = count + 1
			i = i + 1
			j = i - 1
			print items