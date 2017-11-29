def getAllFileNamesInCurrentDir
	allFileNames = Dir::entries(".")
	puts  allFileNames 
	puts

	#ls$B%3%^%s%I$HF1$87k2L$,F@$i$l$F$$$k$+%F%9%H(B
	#ls$B%3%^%s%I$N7k2L$K$O(Bhogehoge.rb$B$,4^$^$l$k$H2>Dj(B
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
	
	#allFileNames$B$,A4$F%U%!%$%k$+%F%9%H(B
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

	#$B3F%U%!%$%k$,E,@Z$JG[Ns$K3JG<$5$l$F$$$k$+%F%9%H(B
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
