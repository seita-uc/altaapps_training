fileEntries = Dir::entries(".")

p fileEntries 

onlyFiles = nil

#$B%U%!%$%k%(%s%H%j$+$i%U%!%$%k$N$_$rCj=P(B
fileEntries.each do |x|
  onlyFiles = fileEntries.delete_if{|x| File::ftype(x) == "directory"}
end

p onlyFiles
