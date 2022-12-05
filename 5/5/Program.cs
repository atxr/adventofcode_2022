using System;
using System.Collections.Generic;
using System.Linq;

namespace _5
{
    class Program
    {
        static Stack<char>[] initStack(string[] inputStack)
        {
            string[] buf = inputStack.Reverse().ToArray();
            int n = Int32.Parse(buf[0].Reverse().SkipWhile(c => c == ' ').First().ToString());
            string[] tmp = new string[n];
            foreach (string line in buf.Skip(1))
            {
                for (int i=0; i<n; i++)
                {
                    char c = line[1 + i * 4];
                    if (c != ' ') {
                        tmp[i] += c;
                    }
                }
            }

            Stack<char>[] stacks = new Stack<char>[n];
            for (int i = 0; i < n; i++)
            {
                stacks[i] = new Stack<char>(tmp[i].ToCharArray());
            }
            return stacks;
        }

        static void execLine(string line, Stack<char>[] stacks, bool mode9001)
        {
            string[] splitedLine = line.Split(' ');
            int k = Int32.Parse(splitedLine[1]);
            int from = Int32.Parse(splitedLine[3]) - 1;
            int to = Int32.Parse(splitedLine[5]) - 1;

            Stack<char> tmpStack;
            if (mode9001)
            {
                tmpStack = new Stack<char>();
                for (int i = 0; i < k; i++)
                {
                    tmpStack.Push(stacks[from].Pop());
                }
            }
            else
            {
                tmpStack = stacks[from];
            }

            for (int i = 0; i<k; i++)
            {
                stacks[to].Push(tmpStack.Pop());
            }
        }

        static void Main()
        { 
            string[] lines = System.IO.File.ReadAllLines(@"./input");
            int sizeHeader = Array.IndexOf(lines, "");
            Stack<char>[] stacks9000 = initStack(lines.Take(sizeHeader).ToArray());
            Stack<char>[] stacks9001 = initStack(lines.Take(sizeHeader).ToArray());

            foreach (string line in lines.Skip(sizeHeader + 1))
            {
                execLine(line, stacks9000, false);
                execLine(line, stacks9001, true);
            }

            string result9000 = "";
            string result9001 = "";
            for (int i = 0; i < stacks9000.Length; i++)
            {
                result9000 += stacks9000[i].Pop();
                result9001 += stacks9001[i].Pop();
            }

            System.Console.WriteLine("Result for mode 9000: {0}", result9000);
            System.Console.WriteLine("Result for mode 9001: {0}", result9001);
            System.Console.ReadLine();
        }
    }
}
