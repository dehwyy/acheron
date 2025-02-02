import clsx from "clsx"
import s from "./ButtonFancy.module.scss"

interface Props extends React.ButtonHTMLAttributes<HTMLButtonElement> {}

export function ButtonFancy({ className, children, ...props }: Props) {
    return (
        <button {...props} className={clsx(className, s.button)}>
            <span className={s.topKey} />
            <p className={s.buttonText}>
                {children}
            </p>
            <span className={s.bottomKey1} />
            <span className={s.bottomKey2} />
        </button>
    )
}
