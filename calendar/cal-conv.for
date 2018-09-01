	program CAL_CONV

	character dread*15,week(7)*9,rec*8
	data week/'Monday','Tuesday','Wednesday','Thursday','Friday',
     *'Saturday','Sunday'/

	print'(19x,5(a/),a)','CALENDAR CONVERSION PROGRAM'
     *,' Good for Julian Day numbers (D) and Gregorian (G), Julian (J) a
     *nd                               Persian or Jalaali (P) calendars'
     *,' ','          v. 1.1 (19 August 2002) [Author: K.M. Borkowski]'
     *,' ',' Today is:'

c Get today's Gregorian date. For older compilers correct possible y2k bug
	call date(dread)				! MM/DD/YY
	if(ichar(dread(7:7)).gt.57) dread(7:7)=char(ichar(dread(7:7))-10)
	read(dread,'(2(i2,1x),i2)') igm,igd,igy
	igy=igy+2000
c Comment out the above 4 lines when using the 'getdat' function below 
C	call getdat(igy,igm,igd)

	JDN=JG2JD(igy,igm,igd,0)
	dread='d'
	go to 25

10    print'(/5x,a)', ' Give one of the four letters DGJP (corresponding
     * to Julian Day Number,     Gregorian, Julian and Persian calendar) 
     * followed by one or three integers of    Julian Day number or date
     * [year,month,day] (e.g. `d2450249'' or `g1996,6,14''):       '

	read(*,'(a)') dread
	if(dread.eq.' ') go to 10

	do 20 i=1,15
	if(dread(i:i).ne.' ') go to 21
20	continue

21	if(dread(i:i).eq.'q'.or.dread(i:i).eq.'Q') stop
	if(dread(i:i).eq.'d'.or.dread(i:i).eq.'D') then
 	read(dread(i+1:15),'(i14)') JDN 	! Julian Day
	else
c	read(dread(i+1:15),'(i5,2i3)') Jy,jm,jd
	read(dread(i+1:15),*) Jy,jm,jd
	endif
	if(dread(i:i).eq.'p'.or.dread(i:i).eq.'P') JDN=Jal2JD(Jy,jm,jd) !Persian
	if(dread(i:i).eq.'g'.or.dread(i:i).eq.'G') JDN=JG2JD(Jy,jm,jd,0)!Gregorian
	if(dread(i:i).eq.'j'.or.dread(i:i).eq.'J') JDN=JG2JD(Jy,jm,jd,1)!Julian

25	call JD2Jal(JDN,jy,jm,jd)
	call JD2JG(JDN,ijy,ijm,ijd,1)
	call JD2JG(JDN,igy,igm,igd,0)

	idw=1+mod(JDN,7)
 	print'(1h ,a9,1x,a,i9,3(3x,a,i5,2i3.2))', week(idw), ' D:',
     * jdn,' G:',igy,igm,igd,' J:',ijy,ijm,ijd, ' P:',jy,jm,jd

	go to 10
         end

       subroutine JalCal(Jy,leap,Gy,March)
c This procedure determines if the Jalaali (Persian) year is 
c leap (366-day long) or is the common year (365 days), and 
c finds the day in March (Gregorian calendar) of the first 
c day of the Jalaali year (Jy)
c Input:  Jy - Jalaali calendar year (-61 to 3177)
c Output:
c   leap  - number of years since the last leap year (0 to 4)
c   Gy    - Gregorian year of the beginning of Jalaali year
c   March - the March day of Farvardin the 1st (1st day of Jy)
       integer breaks(20),Gy
c Jalaali years starting the 33-year rule
       data breaks/-61,9,38,199,426,686,756,818,1111,1181,
     *   1210,1635,2060,2097,2192,2262,2324,2394,2456,3178/
       Gy=Jy+621
       leapJ=-14
       jp=breaks(1)
       if(Jy.lt.jp.or.Jy.ge.breaks(20))  print'(10x,a,i5,a,i5,a)',
     *' Invalid Jalaali year number:',Jy,' (=',Gy,' Gregorian)'
c Find the limiting years for the Jalaali year Jy
       do 1 j=2,20
       jm=breaks(j)
       jump=jm-jp
        if(Jy.lt.jm) go to 2
       leapJ=leapJ+jump/33*8+MOD(jump,33)/4
1      jp=jm
2      N=Jy-jp
c Find the number of leap years from AD 621 to the beginning 
c of the current Jalaali year in the Persian calendar
       leapJ=leapJ+N/33*8+(MOD(N,33)+3)/4
        if(MOD(jump,33).eq.4.and.jump-N.eq.4) leapJ=leapJ+1
c and the same in the Gregorian calendar (until the year Gy)
       leapG=Gy/4-(Gy/100+1)*3/4-150
c Determine the Gregorian date of Farvardin the 1st
       March=20+leapJ-leapG
c Find how many years have passed since the last leap year
       if(jump-N.lt.6) N=N-jump+(jump+4)/33*33
       leap=MOD(MOD(N+1,33)-1,4)
       if(leap.eq.-1) leap=4
	return
          end

       function Jal2JD(Jy,Jm,Jd)
c Converts a date of the Jalaali calendar to the Julian Day Number
c Input:  Jy - Jalaali year (1 to 3100)
c         Jm - month (1 to 12)
c         Jd - day (1 to 29/31)
c Output: Jal2JD - the Julian Day Number
       call JalCal(Jy,leap,iGy,March)
       Jal2JD=JG2JD(iGy,3,March,0)+(Jm-1)*31-Jm/7*(Jm-7)+Jd-1
          end

       subroutine JD2Jal(JDN,Jy,Jm,Jd)
c Converts the Julian Day number to a date in the Jalaali calendar
c Input: JDN - the Julian Day number
c Output: Jy - Jalaali year (1 to 3100)
c         Jm - month (1 to 12)
c         Jd - day (1 to 29/31)

c Calculate Gregorian year (L)
       call JD2JG(JDN,L,M,N,0)
       Jy=L-621
       call JalCal(Jy,leap,iGy,March)
       JDN1F=JG2JD(L,3,March,0)
c Find number of days that passed since 1 Farvardin
       k=JDN-JDN1F
          if(k.ge.0) then
         if(k.le.185) then
c The first 6 months
       Jm=1+k/31
       Jd=MOD(k,31)+1
              return
         else
c The remaining months
       k=k-186
         endif
          else
c previous Jalaali year
       Jy=Jy-1
       k=k+179
       if(leap.eq.1) k=k+1
          endif
       Jm=7+k/30
       Jd=MOD(k,30)+1
          end


       function JG2JD(L,M,N,J1G0)

c Input:  L - calendar year (years BC numbered 0, -1, -2, ...)
c         M - calendar month (for January M=1, February M=2, ..., M=12)
c         N - calendar day of the month M (1 to 28/29/30/31)
c      J1G0 - to be set to 1 for Julian and to 0 for Gregorian calendar
c Output: JG2JD - Julian Day number
c  Calculates the Julian Day number (JG2JD) from Gregorian or Julian
c  calendar dates. This integer number corresponds to the noon of 
c  the date (i.e. 12 hours of Universal Time).
c  The procedure was tested to be good since 1 March, -100100 (of both 
c  the calendars) up to a few millions (10**6) years into the future.
c  The algorithm is based on D.A. Hatcher, Q.Jl.R.Astron.Soc. 25(1984), 53-55
c  slightly modified by me (K.M. Borkowski, Post.Astron. 25(1987), 275-279).

      JG2JD=(L+(M-8)/6+100100)*1461/4+(153*mod(M+9,12)+2)/5+N-34840408
      if(J1G0.LE.0)       JG2JD = JG2JD-(L+100100+(M-8)/6)/100*3/4+752
c     MJD=JG2JD-2400000.5   ! this formula gives Modified Julian Day number
         END


       subroutine JD2JG(JD,L,M,N,J1G0)

c Input:  JD   - Julian Day number
c         J1G0 - to be set to 1 for Julian and to 0 for Gregorian calendar
c Output: L - calendar year (years BC numbered 0, -1, -2, ...)
c         M - calendar month (for January M=1, February M=2, ... M=12)
c         N - calendar day of the month M (1 to 28/29/30/31)
c  Calculates Gregorian and Julian calendar dates from the Julian Day number 
c  (JD) for the period since JD=-34839655 (i.e. the year -100100 of both 
c  the calendars) to some millions (10**6) years ahead of the present.
c  The algorithm is based on D.A. Hatcher, Q.Jl.R.Astron.Soc. 25(1984), 53-55
c  slightly modified by me (K.M. Borkowski, Post.Astron. 25(1987), 275-279).

      J = 4*JD+139361631
      IF(J1G0.LE.0)       J = J+(4*JD+183187720)/146097*3/4*4-3908
      I = MOD(J,1461)/4*5+308
      N = MOD(I,153)/5+1
      M = MOD(I/153,12)+1
      L = J/1461-100100+(8-M)/6
          END
