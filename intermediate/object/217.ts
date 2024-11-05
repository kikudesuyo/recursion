class Files {
  fileName: string;
  fileExtension: string;
  content: string;
  parentFolder: string;

  constructor(
    fileName: string,
    fileExtension: string,
    content: string,
    parentFolder: string
  ) {
    this.fileName = fileName;
    this.fileExtension = fileExtension;
    this.content = content;
    this.parentFolder = parentFolder;
  }

  getLifetimeBandwidthSize(): string {
    if (this.content.length >= 40) {
      return `${String(this.content.length / 40)}GB`;
    }
    return `${String(this.content.length * 25)}MB`;
  }
  prependContent(data: string): string {
    this.content = data + this.content;
    return this.content;
  }

  addContent(data: string, position: number): string {
    return (
      this.content.slice(0, position) + data + this.content.slice(position)
    );
  }
  showFileLocation(): string {
    const FileExtension = [".word", ".png", ".mp4", ".pdf"];
    if (!FileExtension.includes(this.fileExtension)) {
      return `${this.parentFolder} > ${this.fileName}.txt`;
    }
    return `${this.parentFolder} > ${this.fileName}${this.fileExtension}`;
  }
}

const assignment = new Files("assignment", ".word", "ABCDEFGH", "homework");
console.log(assignment.getLifetimeBandwidthSize()); // 200MB
console.log(assignment.prependContent("MMM")); // MMMABCDEFGH
console.log(assignment.addContent("hello world", 5)); // MMMABhello worldCDEFGH
console.log(assignment.showFileLocation()); // homework > assignment.word

const blade = new Files(
  "blade",
  ".txt",
  "bg-primary text-white m-0 p-0 d-flex justify-content-center",
  "view"
);
console.log(blade.getLifetimeBandwidthSize()); // 1.475GB
console.log(blade.addContent("font-weight-bold ", 11)); // ABCDEFGH
console.log(blade.showFileLocation()); // view > blade.txt
