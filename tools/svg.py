import os
import sys


class svg:
    def __init__(self, dir, target):
        self.dir = dir
        self.target = target

    def run(self):
        t = open(self.target, "w")
        t.write("const imports = {\n")

        for dirpath, dirnames, filenames in os.walk(self.dir):
            for filename in filenames:
                file_path = os.path.join(dirpath, filename)
                print("=>", file_path)
                f = os.path.splitext(filename)
                if f[1] != ".svg":
                    print(f)
                    continue
                t.write("'{}' : () => import('{}'),\n".format(f[0], file_path))

            for dirname in dirnames:
                sub_dir_path = os.path.join(dirpath, dirname)
                print("==>", sub_dir_path)
        t.write("};")


e = svg(sys.argv[1], sys.argv[2])
e.run()
