
CurrentDir = Dir::entries(".")

p CurrentDir

FileOnly = CurrentDir

CurrentDir.each do |hoge|
  FileOnly.select!{|hoge| File::ftype(hoge) == "file"}
end

p FileOnly

FileOnly.each do |hoge|
  p hoge
end


