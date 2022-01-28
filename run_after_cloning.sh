cat << EOF >> .git/hooks/pre-commit
make pre_commit
EOF

chmod +x .git/hooks/pre-commit