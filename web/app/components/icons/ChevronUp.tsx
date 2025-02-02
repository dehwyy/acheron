interface IconChevronUpProps {
    className?: string
}

export function IconChevronUp({ className }: IconChevronUpProps) {
    return (
        <svg
            xmlns="http://www.w3.org/2000/svg"
            width="16"
            height="16"
            viewBox="0 0 24 24"
            fill="none"
            stroke="currentColor"
            strokeWidth="2"
            strokeLinecap="round"
            strokeLinejoin="round"
            className={className}
        >
            <path
                stroke="none"
                d="M0 0h24v24H0z"
                fill="none"
            />
            <path d="M6 15l6 -6l6 6" />
        </svg>
    )
}
