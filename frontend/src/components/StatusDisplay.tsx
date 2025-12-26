interface StatusDisplayProps {
  type: "loading" | "error";
  message?: string;
}

const defaultMessages = {
  loading: "Loading...",
  error: "Failed to load configuration",
};

const styles = {
  loading: "text-white",
  error: "text-red-500",
};

export function StatusDisplay({ type, message }: StatusDisplayProps) {
  return (
    <div className="flex items-center justify-center min-h-screen">
      <p className={styles[type]}>{message || defaultMessages[type]}</p>
    </div>
  );
}
