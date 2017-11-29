fileEntries = Dir::entries(".")
puts  fileEntries 
puts

onlyFiles = Array.new

#ファイルエントリからディレクトリを削除
fileEntries.each do |x|
	onlyFiles = fileEntries.delete_if{|x| File::ftype(x) == "directory"}
end

puts  onlyFiles
puts


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

puts "txt file: #{f_txt}"
puts "html file: #{f_html}"
puts "css file: #{f_css}"
puts "md file: #{f_md}"
puts "etc: #{f_etc}"

	
