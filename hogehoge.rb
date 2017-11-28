fileEntries = Dir::entries(".")

p fileEntries 

onlyFiles = nil

#ファイルエントリからファイルのみを抽出
fileEntries.each do |x|
  onlyFiles = fileEntries.delete_if{|x| File::ftype(x) == "directory"}
end

p onlyFiles
