import json

def main():
	input = open('zh-en-booknames.csv')
	zht = []
	zhs = []
	for line in input:
		components = line.rstrip().split('\t')
		assert len(components) == 7
		allNames = set(components[1:5])
		zht.append({
			"ID": components[0],
			"Full": components[1],
			"Abbr": components[2],
			"Aliases": list(sorted(allNames.difference(components[1:3]))),
		})
		zhs.append({
			"ID": components[0],
			"Full": components[3],
			"Abbr": components[4],
			"Aliases": list(sorted(allNames.difference(components[3:5]))),
		})
	zht_out = open("zh-hant.bible.unv/book_names.json", "w")
	zht_out.write(json.dumps(zht, indent=2, ensure_ascii=False))
	zht_out.close()

	zhs_out = open("zh-hans.bible.unv/book_names.json", "w")
	zhs_out.write(json.dumps(zhs, indent=2, ensure_ascii=False))
	zhs_out.close()

if __name__ == '__main__':
	main()
