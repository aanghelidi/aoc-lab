nums = IO.readlines('input.txt').map(&:to_i)
# Part 1
puts "Part 1: #{nums.reduce(0) { |sum, line| sum + line }}"
# Part 2
frequencies = Set.new
frequency = 0
nums.cycle do |line|
  frequency += line
  if frequencies.include?(frequency)
    puts "Part 2: #{frequency}"
    break
  end
  frequencies.add(frequency)
end
