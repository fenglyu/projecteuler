#include <iostream>
#include <algorithm>    /// next_permutation, sort
#include <string>
#include <functional>

using namespace std;

int main () {
  int myints[] = {1,2,3,4,5};
  sort (myints,myints+4);

  cout << "After sort: " << myints[0] << ' ' << myints[1] << ' ' << myints[2] << ' '<< myints[3] << ' ' << myints[4] << '\n';
  do {
    cout << myints[0] << ' ' << myints[1] << ' ' << myints[2] << ' '<< myints[3]<< ' ' << myints[4] <<'\n';
  } while ( next_permutation(myints,myints+5) );    ///获取下一个较大字典序排列

  cout << "After loop: " << myints[0] << ' ' << myints[1] << ' ' << myints[2] << ' '<< myints[3] << ' ' << myints[4] <<'\n';


/*
  string s="abc";
  cout << "aaa:" << s.begin() << ' ' << s.end() << '\n';
  sort(s.begin(), s.end(), greater<char>());

  do {
     cout << s << ' ';
  } while(prev_permutation(s.begin(), s.end()));
  cout << '\n';
*/
  return 0;
}
