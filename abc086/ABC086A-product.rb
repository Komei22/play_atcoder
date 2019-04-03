a, b = gets.chomp.split(" ").map(&:to_i)

c = a*b

c.odd? ? print("Odd") : print("Evenj")
# if c.odd?
#     print("Even")
# else
#     print("Odd")
# end
