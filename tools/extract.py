import os
import re
import sys
from os import path

exclude_patterns = ["{{ template .*}}", "{{ svg .*}}"]

variable_patterns = [
    "{{ ?(if )?(.i18n.Tr )?(.*) ?}}",
]

r"[^0-9\s]+"
r"\b(\w+)\b(?!.*\b\1\b)"


class extract:
    def __init__(self, src, dst):
        self.src = src
        self.dst = dst

    def run(self):
        t = open(self.dst, "w")

        with open(self.src, "r") as f:
            line = f.readline()
            while line:
                result = self.extract(line)
                if result != "":
                    t.write(result)
                    t.write("\n")
                line = f.readline()

    def extract(self, line):
        r = re.search(
            r"\{\{.*\}\}",
            line,
        )
        if r != None:
            # print(r[0])
            t = r[0]
            for ep in exclude_patterns:
                t = re.sub(
                    r"^(.*?){}(.*?)$".format(ep),
                    lambda x: x.group(1) + " " * len(x.group(2)),
                    t,
                )
            # print(t)
            if len(t) <= 0:
                return ""

            # r"{}".format(vp),
            # r"\b(\w+)\b(?!.*\b\1\b)",
            # "{{ ?(if )?(.i18n.Tr )?(.*) ?}}",
            # for vp in variable_patterns:
            v = re.search(
                r"{{ ?(if )?(.i18n.Tr )?(.*) ?}}",
                t,
            )
            print("1:", v.group(1), "2:", v.group(2), ",3:", v.group(3))
            return t
        return ""


e = extract(sys.argv[1], sys.argv[2])
e.run()
