#include <stdio.h>
#include <curses.h>
#include "main.h"

int isPrime(int x) {
	if (x < 2) {
		return 0;
	} else if (x == 2) {
		return 1;
	} else {
		for (int i = 2; i <= x/2; ++i) {
			if (x % i == 0) {
				return 0;
			}
		}
	}
	return 1;
}
