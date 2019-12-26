set yrange [0.0:1.0]
set xlabel "S" font ",10"
set ylabel "Cumulative Probability" font ",10"
set grid ytics lt 0 lw 1 lc rgb "#bbbbbb"
set grid xtics lt 0 lw 1 lc rgb "#bbbbbb"


plot "distr_S_T_reflected.data" using 1:2 title "original" with lines lt rgb "#666666",\
   "distr_S_T_reflected.data" using 1:3 title "reflected" with lines lt rgb "#00FF7F"