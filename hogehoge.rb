# o = Dir::entries(".")
# i = Dir::glob("*/")
# s = nil
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

# # i.each do |hoge|
# #   s << hoge.delete("/")
# # end

# # p s

# # i.each do |hoge|
# #   p hoge.to_s
# # end

# # o.each do |hoge|
# #   p hoge.to_s
# # end

# o.each do |del|
#   o.delete(del) if i.include?(del)
# end

# puts o
# puts "\n"
# puts i


