require 'test/unit'

class Test_Sample < Test::Unit::TestCase
  def setup
	@obj = Test_Sample.new
  end

  # def teardown
  # end

  def test_getAllFileNamesInCurrentDir
	 assert_equal("foo", @obj.getAllFileNamesInCurrentDir)
  end
end
