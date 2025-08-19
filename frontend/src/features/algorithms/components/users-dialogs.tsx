import { useAlgorithms } from '../context/algorithms-context';
import { UsersActionDialog } from './users-action-dialog';

export function UsersDialogs() {
  const { open, setOpen, currentRow, setCurrentRow } = useAlgorithms();
  return (
    <>
      {currentRow && (
        <>
          <UsersActionDialog
            key={`user-edit-${currentRow.id}`}
            open={open === 'edit'}
            onOpenChange={() => {
              setOpen('edit');
              setTimeout(() => {
                setCurrentRow(null);
              }, 500);
            }}
            currentRow={currentRow}
          />
        </>
      )}
    </>
  );
}
