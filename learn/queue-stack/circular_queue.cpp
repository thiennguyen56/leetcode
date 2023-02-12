#include <iostream>
#include <vector>

using namespace std;

class MyCircularQueue {
private:
  vector<int> data;
  int length;
  int first = 0;
  int last = 0;

public:
  MyCircularQueue(int k) { length = 5; }

  bool enQueue(int value) {}
};
