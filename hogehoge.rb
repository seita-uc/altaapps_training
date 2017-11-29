def getAllFileNamesInCurrentDir
	allFileNames = Dir::entries(".")
	puts  allFileNames 
	puts

	#lsコマンドと同じ結果が得られているかテスト
	#lsコマンドの結果にはhogehoge.rbが含まれると仮定
	if allFileNames.include?("hogehoge.rb")
		puts true
	else
		puts false
	end

	return allFileNames
end

def extractOnlyFiles(allFileNames)
	onlyFiles = Array.new

	allFileNames.each do |x|
		onlyFiles = allFileNames.delete_if{|x| File::ftype(x) == "directory"}
	end

	puts onlyFiles
	puts
	
	#allFileNamesが全てファイルかテスト
	puts allFileNames.all? {|x| File::ftype(x) == "file"}	

	return onlyFiles	
end

def assortFileTypes(onlyFiles)
	f_txt = Array.new 
	f_html = Array.new 
	f_css = Array.new 
	f_md = Array.new 
	f_etc = Array.new 

	onlyFiles.each do |x|
		if    File.extname(x) == ".txt"
			f_txt.push(x) 
		elsif File.extname(x) == ".html"		
			f_html.push(x) 
		elsif File.extname(x) == ".css"		
			f_css.push(x)		
		elsif File.extname(x) == ".md"				
			f_md.push(x)		
		else 
			f_etc.push(x)	
		end
	end
	
	puts "txt file: "
	puts f_txt
	puts "html file: "
	puts f_html
	puts "css file: "
	puts f_css
	puts "md file: "
	puts f_md
	puts "etc: "
	puts f_etc

	#各ファイルが適切な配列に格納されているかテスト
	if f_txt.include?("foo.txt")
		puts true
	else
	 	puts false
	end

	if f_html.include?("fooo.html")
		puts true
	else
	 	puts false
	end

	if f_css.include?("foo.css")
		puts true
	else
	 	puts false
	end

	if f_md.include?("README.md")
		puts true
	else
	 	puts false
	end

	if f_etc.include?("hogehoge.rb")
		puts true 
	else
	 	puts false
	end
end

fileEntries = Array.new
fileEntries  = getAllFileNamesInCurrentDir
onlyFiles = Array.new
onlyFiles = extractOnlyFiles(fileEntries)
assortFileTypes(onlyFiles)
