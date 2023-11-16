lines = IO.readlines("input.txt").map(&:chomp)

def check_counts(counts, n) = counts.values.include?(n)
def count_diff(line1, line2) = line1.each_char.with_index.count { |char, i| char != line2[i] }
def find_common_letter(line1, line2) = line1.each_char.with_index.map { |char, i| char if char == line2[i] }.join

countains_two = 0
countains_three = 0
lines.each_with_index do |line, i|
  counts = Hash.new(0)
  line.each_char { |char| counts[char] += 1 }
  countains_two += 1 if check_counts(counts, 2)
  countains_three += 1 if check_counts(counts, 3)
  # Part 2
  lines.slice(0..i-1).each { |line2| puts "Part 2: #{find_common_letter(line, line2)}" if count_diff(line, line2) == 1 }
end
puts "Part 1: #{countains_two * countains_three}"
