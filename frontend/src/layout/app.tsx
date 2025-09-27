export const AppLayout = ({ children }: { children: React.ReactNode }) => {
  return (
    <div className="wrapper min-h-screen flex relative lg:static surface-ground">
      {children}
    </div>
  );
};
