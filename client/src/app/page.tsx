import SudokuGame from "./components/SudokuGame";

export default function Home() {
  return (
    <main className="flex min-h-screen flex-col items-center p-16 space-y-4">
      <h1>The Sudoku solver</h1>
      <div>
        <SudokuGame />
      </div>
    </main>
  );
}
