import { Notification } from '@carbon/icons-react';
import { HeaderGlobalAction, ToastNotification } from '@carbon/react';
import { css } from '@emotion/css';
import { useEffect, useRef, useState } from 'react';

interface Props {
  notifications: Notification[];
}

export interface Notification {
  id?: number;
  title: string;
  subtitle?: string;
  type: 'error' | 'info' | 'info-square' | 'success' | 'warning' | 'warning-alt';
  caption?: string;
}

export const Notifications = ({ notifications }: Props): JSX.Element => {
  const [isNotificationsOpen, setIsNotificationsOpen] = useState(false);
  const ref = useRef<HTMLDivElement>(null);

  useEffect(() => {
    const handleClickOutside = (e: MouseEvent) => {
      if (ref.current && !ref.current.contains(e.target as Node)) {
        setIsNotificationsOpen(false);
      }
    };
    document.addEventListener('mousedown', handleClickOutside);
    return () => {
      document.removeEventListener('mousedown', handleClickOutside);
    };
  }, []);

  return (
    <div ref={ref} className={styles.wrapper}>
      <HeaderGlobalAction aria-label="Notifications" onClick={() => setIsNotificationsOpen(!isNotificationsOpen)}>
        <Notification size={20} />
      </HeaderGlobalAction>
      <div className={styles.notifs(isNotificationsOpen)}>
        {notifications.map((notif, idx) => (
          <ToastNotification key={idx} title={notif.title} subtitle={notif.subtitle} caption={notif.caption} onClose={() => false} />
        ))}
      </div>
    </div>
  );
};

const styles = {
  notifs: (isOpen: boolean) => css`
    position: absolute;
    right: 0;
    margin-top: 1rem;
    margin-right: ${isOpen ? '0rem' : '-23rem'};
    width: 23rem;
    transition: margin-right 0.3s;

    display: flex;
    flex-direction: column;
    gap: 1rem;
  `,
  wrapper: css`
    height: calc(100% - 3rem);
  `,
};
