# SWEA9015번 : 배열의 분할

## 문제
나삼성씨는 길이 n의 배열이 주어질 때, 이 배열을 연속된 부분배열들로 적당히 분할해서
 
각 부분 배열이 같거나 증가하는 오름차순이거나 같거나 감소하는 내림차순으로 정렬되어 있기를 원한다.
 
나삼성씨는 부분 배열의 분할 개수를 최소로 하고 싶다. 배열을 최소 몇 개로 분할해야 목적을 이룰 수 있을까?

## 입력
첫 번째 줄에 테스트 케이스의 수 TC가 주어진다.
 
이후 TC개의 테스트 케이스가 새 줄로 구분되어 주어진다.

각 테스트 케이스는 다음과 같이 구성되었다.


첫 번째 줄에 배열의 크기 n이 주어진다. (1 ≤ n ≤ 100,000)
 
이후 배열을 이루는 n개의 정수 ai가 주어진다. (1 ≤ ai ≤ 109)

## 출력
각 테스트 케이스마다 ‘#x’(x는 테스트케이스 번호를 의미하며 1부터 시작한다)를 출력하고,
 

각 테스트 케이스마다 한 줄씩 민규가 분할해서 만들어야 하는 최소 배열의 개수를 출력하라.