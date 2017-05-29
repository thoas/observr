"""initial

Revision ID: bdd74fcd8a5c
Revises: 
Create Date: 2017-05-24 13:32:54.409856

"""
from alembic import op
import sqlalchemy as sa


# revision identifiers, used by Alembic.
revision = 'bdd74fcd8a5c'
down_revision = None
branch_labels = None
depends_on = None

from sqlalchemy.dialects.postgresql import UUID, JSON


def upgrade():
    op.create_table('observr_user',
        sa.Column('id', UUID, primary_key=True),
        sa.Column('username', sa.String(255), unique=True, nullable=False),
        sa.Column('email', sa.String(255), unique=True, nullable=False),
        sa.Column('password', sa.String(255), nullable=False),
        sa.Column('created_at', sa.DateTime(timezone=True), server_default=sa.func.now(), nullable=False),
        sa.Column('updated_at', sa.DateTime(timezone=True), server_default=sa.func.now(), nullable=False),
        sa.Column('deleted_at', sa.DateTime(timezone=True), nullable=True),
    )

    op.create_table('observr_project',
        sa.Column('id', UUID, primary_key=True),
        sa.Column('name', sa.String(255), nullable=False),
        sa.Column('url', sa.String(2048), nullable=False),
        sa.Column('api_key', sa.String(255), unique=True, nullable=False),
        sa.Column('created_at', sa.DateTime(timezone=True), server_default=sa.func.now(), nullable=False),
        sa.Column('updated_at', sa.DateTime(timezone=True), server_default=sa.func.now(), nullable=False),
        sa.Column('deleted_at', sa.DateTime(timezone=True), nullable=True),
        sa.Column('user_id', UUID, sa.ForeignKey('observr_user.id', ondelete="CASCADE"), nullable=False),
    )

    op.create_table('observr_visit',
        sa.Column('id', UUID, primary_key=True),
        sa.Column('host', sa.String(255), nullable=False),
        sa.Column('path', sa.String(2048), nullable=True),
        sa.Column('referer', sa.String(255), nullable=True),
        sa.Column('remote_addr', sa.String(255), nullable=False),
        sa.Column('method', sa.String(10), nullable=False),
        sa.Column('user_agent', sa.String(255), nullable=True),
        sa.Column('status_code', sa.Integer, nullable=True),
        sa.Column('protocol', sa.String(20), nullable=True),
        sa.Column('data', JSON(), nullable=True),
        sa.Column('headers', JSON(), nullable=True),
        sa.Column('query_string', sa.String(255), nullable=True),
        sa.Column('created_at', sa.DateTime(timezone=True), server_default=sa.func.now(), nullable=False),
        sa.Column('updated_at', sa.DateTime(timezone=True), server_default=sa.func.now(), nullable=False),
        sa.Column('project_id', UUID, sa.ForeignKey('observr_project.id', ondelete="CASCADE"), nullable=False),
    )

    op.create_table('observr_tag',
        sa.Column('id', UUID, primary_key=True),
        sa.Column('key', sa.String(255), unique=True, nullable=False),
        sa.Column('value', sa.Text, nullable=False),
        sa.Column('data', sa.Text, nullable=False),
        sa.Column('seen_count', sa.Integer, nullable=False, server_default="0"),
        sa.Column('project_id', UUID, sa.ForeignKey('observr_project.id', ondelete="CASCADE"), nullable=False),
        sa.Column('created_at', sa.DateTime(timezone=True), server_default=sa.func.now(), nullable=False),
        sa.Column('updated_at', sa.DateTime(timezone=True), server_default=sa.func.now(), nullable=False),
        sa.Column('first_seen_at', sa.DateTime(timezone=True), nullable=False),
        sa.Column('last_seen_at', sa.DateTime(timezone=True), nullable=False),
    )

    op.create_table('observr_group_tag',
        sa.Column('id', UUID, primary_key=True),
        sa.Column('created_at', sa.DateTime(timezone=True), server_default=sa.func.now(), nullable=False),
        sa.Column('src_tag_id', UUID, sa.ForeignKey('observr_tag.id', ondelete="CASCADE"), nullable=False),
        sa.Column('dst_tag_id', UUID, sa.ForeignKey('observr_tag.id', ondelete="CASCADE"), nullable=False),
    )

    op.create_table('observr_visit_tag',
        sa.Column('id', UUID, primary_key=True),
        sa.Column('created_at', sa.DateTime(timezone=True), server_default=sa.func.now(), nullable=False),
        sa.Column('visit_id', UUID, sa.ForeignKey('observr_visit.id', ondelete="CASCADE"), nullable=False),
        sa.Column('tag_id', UUID, sa.ForeignKey('observr_tag.id', ondelete="CASCADE"), nullable=False),
    )


def downgrade():
    op.drop_table('observr_group_tag')
    op.drop_table('observr_visit_tag')
    op.drop_table('observr_tag')
    op.drop_table('observr_visit')
    op.drop_table('observr_project')
    op.drop_table('observr_user')
