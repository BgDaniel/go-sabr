set yrange [0.0:1.0]
set xlabel "Time" font ",10"
set ylabel "Cumulative Probability" font ",10"
set grid ytics lt 0 lw 1 lc rgb "#bbbbbb"
set grid xtics lt 0 lw 1 lc rgb "#bbbbbb"


plot "absor_distr.data" using 1:2 title "implicite" with lines lt rgb "#666666",\
   "absor_distr.data" using 1:3 title "explicite" with lines lt rgb "#00FF7F"